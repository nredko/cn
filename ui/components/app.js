import Vue from '../lib/vue.esm.browser.min.js'
import {NavbarComponent} from "./navbar.js";
import {ContainerComponent} from "./container.js";

new Vue({
  el: '#app',
  components: {
    navbar: NavbarComponent,
    container: ContainerComponent,
  },
  template: `
<div>
    <navbar v-on:search='search'></navbar>
    <container ref='container' style='margin-top: 50px;'></container>
</div>`,
  methods: {
    search: function (text) {
      this.$refs.container.search(text);
    }
  }
});
