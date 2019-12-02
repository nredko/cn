import {NotarizationService} from "../service/notarization.js";

const NotarizationButtonComponent = {
  template: `
<div>
  <button
    v-if='!isNotarized() || isUnknown()'
  	role='button'
  	type='button'
  	class='btn btn-outline-primary'
  	v-on:click='notarize("Notarized")'>
  	Notarize
  </button>
  <button
    v-if='isNotarized() || isUnknown()'
  	role='button'
  	type='button'
  	class='btn btn-outline-danger'
  	v-on:click='notarize("Untrusted")'>
  	Un-trust
  </button>
</div>`,
  props: ['image'],
  data: function () {
    return {
      image: {},
    }
  },
  methods: {
    isNotarized: function () {
      return this.image.Notarization.Status === 'Notarized';
    },
    isUnknown: function () {
      return !this.image.Notarization.Status || this.image.Notarization.Status === 'Unknown';
    },
    notarize: function (status) {
      const self = this;
      NotarizationService.notarize(this.image.Image.Hash, status, function () {
        self.$emit("refresh");
      });
    },
  }
};

export {NotarizationButtonComponent};
