import {NotarizationService} from "../service/notarization.js";

const BulkNotarizationComponent = {
  props: ['filter'],
  template: `
<div>
	<button
	  type='button'
    role='button'
		class='btn btn-outline-primary'
		v-on:click='bulkNotarize()'>
		Bulk Notarize
	</button>
	<button
	  type='button'
		role='button'
		class='btn btn-outline-danger'
		v-on:click='bulkUnTrust()'>
		Bulk Un-trust
	</button>
</div>`,
  data: function () {
    return {
      filter: '',
    }
  },
  methods: {
    bulkNotarize: function () {
      const self = this;
      NotarizationService.bulkNotarize(self.filter, function () {
        self.$emit('refresh');
      });
    },
    bulkUnTrust: function () {
      const self = this;
      NotarizationService.bulkUnTrust(self.filter, function () {
        self.$emit('refresh');
      });
    }
  }
};

export {BulkNotarizationComponent};
