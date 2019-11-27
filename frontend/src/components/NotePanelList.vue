<template>
  <div class="note-panel-list">
    <div class="header">
      <h2 class="title">Manage Your Notes</h2>
    </div>
    <div class="body">
      <div v-if="notes.length > 0" class="note-container">
        <div v-for="note in notes" :key="note.id" @click="$emit('note-click', note.id)" class="note">
          <div class="info">
            <span class="note">{{ note.note }}</span>
            <span class="date">Modified: {{ note.modifiedDate | beautifyDate }}</span>
          </div>
        </div>
      </div>
      <div v-else class="empty-container">
        Press the pluss at the bottom to add a new note
      </div>
    </div>
    <div class="footer">
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
      notes: []
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
        width: 100%;

        & > .info {
          bottom: 0;
          cursor: pointer;
          left: 0;
          position: absolute;
          right: 0;
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