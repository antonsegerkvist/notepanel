<template>
  <div class="note-panel-list">
    <div class="header">
      <h2 class="title">Manage Your Notes</h2>
    </div>
    <div class="body">
      <div v-if="notes.length > 0" class="note-container">
        <div v-for="note in notes" :key="note.id" class="note">
          <div @click="$emit('note-click', note.id)" class="info">
            <span class="note">{{ note.note }}</span>
            <span v-if="note.modifiedDate !== note.createDate" class="date">
              Modified: {{ note.modifiedDate | beautifyDate }}
            </span>
            <span v-else class="date">
              Created: {{ note.modifiedDate | beautifyDate }}
            </span>
          </div>
          <div class="select-area">
            <label class="label-checkbox">
              <input
                @click="event => handleCheckboxClick(event.target.checked, note.id)"
                type="checkbox"
                class="checkbox">
              <span class="checkmark"></span>
            </label>
          </div>
        </div>
      </div>
      <div v-else class="empty-container">
        Press the pluss at the bottom to add a new note
      </div>
    </div>
    <div class="footer">
      <div :class="deleteButtonClass" @click="handleDeleteClick">
        <i class="material-icons icon">delete</i>
      </div>
      <div @click="$emit('add-note-click')" class="button button-add">
        + Add Note
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({

  name: 'note-panel-list',

  data () {
    return {
      keyword: '',
      notes: [],
      markedNotes: []
    }
  },

  created () {
    this.fetchNotes();
  },

  filters: {
  
    beautifyDate (dateString) {
      let date = new Date(dateString);

      let year = date.getFullYear();
      let month = date.getMonth() + 1;
      let day = (date.getDate() + 1);
      let hour = date.getHours() + '';
      let minute = date.getMinutes() + '';

      if (hour.length < 2) hour = '0' + hour;
      if (minute.length < 2) minute = '0' + minute;

      month = ({
        '1': 'Jan',
        '2': 'Feb',
        '3': 'Mar',
        '4': 'April',
        '5': 'May',
        '6': 'June',
        '7': 'July',
        '8': 'Aug',
        '9': 'Sep',
        '10': 'Oct',
        '11': 'Nov',
        '12': 'Dec'
      })[month];

      return `${day} ${month} ${year} ${hour}:${minute}`;

    }
  
  },

  methods: {

    fetchNotes () {
      var xhr = new XMLHttpRequest();

      xhr.onreadystatechange = () => {
        if (xhr.readyState === 4)
        {
          if (xhr.status === 200)
          {
            let o = JSON.parse(xhr.responseText);
            if (o)
            {
              this.notes = o.notes;
            }
          }
        }
      };

      xhr.open('GET', '/api/note/list');
      xhr.send();
    },

    handleCheckboxClick (checked, noteID) {
      let index = this.markedNotes.indexOf(noteID);
      if (checked === true && index === -1)
      {
        this.markedNotes.push(noteID);
      }
      else if (checked === false && index !== -1)
      {
        this.markedNotes.splice(index, 1);
      }
    },

    handleDeleteClick () {
      if (this.markedNotes.length === 0)
      {
        return;
      }

      var xhr = new XMLHttpRequest();

      xhr.onreadystatechange = () => {
        if (xhr.readyState === 4)
        {
          if (xhr.status === 202)
          {
            this.fetchNotes();
          }
        }
      };

      xhr.open('DELETE', '/api/note');
      xhr.setRequestHeader('Content-Type', 'application/json');
      xhr.send(JSON.stringify({
        ids: this.markedNotes
      }));
      this.markedNotes = [];
    }

  },

  computed: {

    deleteButtonClass () {
      let ret = ['small-button', 'small-button-delete'];
      if (this.markedNotes.length === 0)
      {
        ret.push('small-button-inactive');
      }
      return ret;
    }

  }

});
</script>

<style lang="scss" scoped>
.note-panel-list {
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
    bottom: 50px;
    left: 0;
    position: absolute;
    right: 0;
    top: 51px;

    & > .note-container {
      bottom: 0;
      box-sizing: border-box;
      left: 0;
      overflow-x: hidden;
      overflow-y: auto;
      padding: 10px;
      position: absolute;
      right: 0;
      top: 0;

      & > .note {
        background-color: #eee;
        float: left;
        height: 50px;
        margin-bottom: 10px;
        position: relative;
        transition: background-color 100ms ease-in-out;
        width: 100%;

        &:hover {
          background-color: #f5f5f5;
        }

        & > .info {
          bottom: 0;
          cursor: pointer;
          left: 0;
          position: absolute;
          right: 25px;
          top: 0;

          & > .note,
          & > .date {
            box-sizing: border-box;
            height: 25px;
            left: 0;
            line-height: 25px;
            overflow: hidden;
            padding: 0 10px;
            position: absolute;
            right: 0;
            text-overflow: ellipsis;
            white-space: nowrap;
          }

          & > .note {
            font-size: 16px;
            top: 0;
          }
          & > .date {
            font-size: 14px;
            font-style: italic;
            bottom: 0;
          }
        }

        & > .select-area {
          bottom: 0;
          box-sizing: border-box;
          position: absolute;
          right: 0;
          top: 0;
          width: 25px;

          & > .label-checkbox {
            box-sizing: border-box;
            cursor: pointer;
            height: 15px;
            left: 50%;
            position: absolute;
            top: 50%;
            transform: translate(-50%, -50%);
            width: 15px;

            & > .checkbox {
              height: 1px;
              position: absolute;
              opacity: 0;
              width: 1px;
            }
            
            & > .checkmark {
              background-color: #fff;
              border: 1px solid #333;
              border-radius: 3px;
              box-sizing: border-box;
              display: inline-block;
              height: 100%;
              width: 100%;
            }

            & > .checkbox:checked ~ .checkmark::before {
              border-left: 1px solid #000;
              border-bottom: 1px solid #000;
              content: '';
              display: inline-block;
              height: 3px;
              left: 50%;
              position: absolute;
              top: 50%;
              transform: translate(-50%, -50%) rotate(-45deg);
              width: 5px;
            }
          }
        }
      }
    }

    & > .empty-container {
      font-style: italic;
      left: 50%;
      line-height: 1.5;
      max-width: 100%;
      position: absolute;
      text-align: center;
      top: 50%;
      transform: translate(-50%, -50%);
      width: 150px;
    }
  }

  & > .footer {
    border-top: 1px solid #eee;
    bottom: 0;
    height: 50px;
    left: 10px;
    position: absolute;
    right: 10px;

    & > .small-button {
      background-color: #333;
      color: #fff;
      cursor: pointer;
      height: 35px;
      line-height: 35px;
      position: absolute;
      text-align: center;
      top: 50%;
      transform: translateY(-50%);
      width: 35px;

      &-inactive {
        opacity: 0.5;
      }

      & > .icon {
        font-size: 16px;
        left: 50%;
        position: absolute;
        top: 50%;
        transform: translate(-50%, -50%);
      }
    }

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
    
      &-add {
        right: 0;
      }

      &:hover {
        background-color: #555;
      }
    }
  }
}
</style>