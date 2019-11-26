<template>
  <div @click="$emit('toggle')" :class="getClass">
    <i class="material-icons icon">note</i>
    <span class="title">Notes</span>
  </div>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({

  name: 'toggle',

  props: {

    open: {
      type: Boolean,
      default () { return false; }
    }

  },

  computed: {

    getClass () {
      let ret = ['toggle'];
      if (this.open === true)
      {
        ret.push('toggle-open');
      }
      return ret;
    }

  }

});
</script>

<style lang="scss" scoped>
.toggle {
  border-top-left-radius: 5px;
  border-bottom-left-radius: 5px;
  bottom: 20px;
  box-shadow: -5px 0 5px rgba(0, 0, 0, 0.2);
  cursor: pointer; 
  font-family: 'Source Sans Pro', sans-serif;
  height: 150px;
  position: fixed;
  right: 0;
  transition: right 100ms ease-in-out;
  user-select: none;
  width: 35px;

  &-open {
    right: 300px;
  }

  & > .icon {
    font-size: 18px;
    left: 50%;
    position: absolute;
    transform: translateX(-50%);
    top: 10px;
  }

  & > .title {
    left: 50%;
    position: absolute;
    top: 50%;
    transform: translate(-50%, -50%) rotate(-90deg);
  }
}
</style>