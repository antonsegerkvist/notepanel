<template>
  <div :class="getClass">
    <NotePanelList
      v-if="currentView === Views.ListView"
      @note-click="handleNoteClick"
      @add-note-click="handleNoteAddClick"/>
    <NotePanelEditor
      v-else-if="currentView === Views.EditView"
      :id="selectedNoteID"/>
  </div>
</template>

<script>
import Vue from 'vue';
import NotePanelList from './NotePanelList.vue';
import NotePanelEditor from './NotePanelEditor.vue';

export default Vue.extend({

  name: 'note-panel',

  components: {
    NotePanelList,
    NotePanelEditor
  },

  props: {

    open: {
      type: Boolean,
      default () { return false; }
    }

  },

  data () {
    return {
      Views: {
        ListView: 0,
        EditView: 1
      },

      currentView: 0,
      selectedNoteID: 0
    };
  },

  methods: {

    handleNoteClick (noteID) {
      this.selectedNoteID = noteID;
      this.currentView = this.Views.EditView;
    },

    handleNoteAddClick () {
      this.selectedNoteID = 0;
      this.currentView = this.Views.EditView;
    }

  },

  computed: {

    getClass () {
      let ret = ['note-panel'];
      if (this.open === true)
      {
        ret.push('note-panel-open');
      }
      return ret;
    }

  }

});
</script>

<style lang="scss" scoped>
.note-panel {
  background-color: #fff;
  bottom: 0;
  box-shadow: 0 0 5px rgba(0,0,0,0.2);
  position: fixed;
  right: -300px;
  top: 0;
  transition: right 100ms ease-in-out;
  width: 300px;

  &-open {
    right: 0;
  }
}
</style>