<template>
  <div class="note-panel-editor">
    <div class="header">
      <h2 class="title">Edit Note</h2>
    </div>
    <div class="body">
      <div v-if="errorType === 1" class="error-window">
        Could not fetch the note information<br>
        (Error Code: {{ error }})
      </div>
      <div v-else-if="errorType === 2" class="error-window">
        Could not save the note information<br>
        (Error Code: {{ error }})
      </div>
      <textarea v-else v-model="note" placeholder="Edit your note here" class="editor">
      </textarea>
    </div>
    <div class="footer">
      <div @click="$emit('back-click')" class="button button-back">
        Back
      </div>
      <div @click="handleSaveNote" class="button button-save">
        Save
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({

  name: 'note-panel-editor',

  props: {

    id: {
      type: Number,
      default () { return 0; }
    }

  },

  data () {
    return {
      note: '',
      createDate: '',
      modifiedDate: '',

      errorType: 0,
      error: 0
    }
  },

  created () {
    if (this.id !== 0)
    {
      this.fetchNote();
    }
  },

  methods: {

    fetchNote () {
      let xhr = new XMLHttpRequest();

      xhr.onreadystatechange = () => {
        if (xhr.readyState === 4)
        {
          if (xhr.status === 200)
          {
            let o = JSON.parse(xhr.responseText);
            if (o)
            {
              this.note = o.note.note;
              this.createDate = o.note.createDate;
              this.modifiedDate = o.note.modifiedDate;
            }
          }
          else
          {
            this.error = xhr.status;
            this.errorType = 1;
          }
        }
      };

      xhr.open('GET', `/api/note/${this.id}`);
      xhr.send();
    },

    handleSaveNote () {
      let method = '';
      let endpoint = '';

      if (this.id !== 0)
      {
        method = 'PATCH';
        endpoint = `/api/note/${this.id}`;
      }
      else
      {
        method = 'POST';
        endpoint = '/api/note';
      }
      this.saveNote(method, endpoint);
    },

    saveNote (method, endpoint) {
      let xhr = new XMLHttpRequest();

      xhr.onreadystatechange = () => {
        if (xhr.readyState === 4)
        {

          if (xhr.status === 202)
          {
            this.$emit('note-updated');
          }
          else if (xhr.status === 201)
          {
            let o = JSON.parse(xhr.responseText);
            this.$emit('note-created', o.id)
          }
          else
          {
            this.error = xhr.status;
            this.errorType = 2;
          }
        }
      };

      xhr.open(method, endpoint);
      xhr.setRequestHeader('Content-Type', 'application/json');
      xhr.send(JSON.stringify({
        note: this.note
      }));
    }

  }

});
</script>

<style lang="scss" scoped>
.note-panel-editor {
  bottom: 0;
  font-family: 'Source Sans Pro', sans-serif;
  left: 0;
  position: absolute;
  right: 0;
  top: 0;

  & > .header {
    border-bottom: 1px solid #eee;
    height: 50px;
    left: 10px;
    line-height: 50px;
    font-size: 18px;
    position: absolute;
    right: 10px;
    top: 0;
  }

  & > .body {
    bottom: 61px;
    left: 10px;
    position: absolute;
    right: 10px;
    top: 61px;

    & > .error-window {
      font-style: italic;
      left: 50%;
      line-height: 1.5;
      position: absolute;
      text-align: center;
      top: 50%;
      transform: translate(-50%, -50%);
    }

    & > .editor {
      appearance: none;
      border: 1px solid #fff;
      box-sizing: border-box;
      font-family: 'Source Sans Pro', sans-serif;
      font-size: 16px;
      height: 100%;
      margin: 0;
      padding: 0;
      resize: none;
      transition: border-color 100ms ease-in-out;
      width: 100%;

      &:hover {
        border-color: #eee;
      }
    }
  }

  & > .footer {
    bottom: 0;
    border-top: 1px solid #eee;
    height: 50px;
    left: 10px;
    position: absolute;
    right: 10px;

    & > .button {
      background-color: #333;
      color: #fff;
      cursor: pointer;
      height: 35px;
      line-height: 35px;
      padding: 0 20px;
      position: absolute;
      top: 50%;
      transform: translateY(-50%);

      &-back { left: 0; }
      &-save { right: 0; }

      &:hover {
        background-color: #555;
      }
    }
  }
}
</style>