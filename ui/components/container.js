import {NotarizationService} from "../service/notarization.js";
import {LoadingComponent} from "./loading.js";
import {NotarizationTableComponent} from "./notarization-table.js";
import {BulkNotarizationComponent} from "./bulk-notarization.js";
import {ConfigurationService} from "../service/configuration.js";

const ContainerComponent = {
  template: `
<div>
  <loading 
    v-if='!initialized'>
  </loading> 
  <div 
    v-if='initialized' 
    class='container'>
	  <notarizationTable 
	    v-bind:images='images' 
	    v-on:refresh='refresh'>
    </notarizationTable>
    <bulkNotarization 
      v-bind:filter='filter'
      v-on:refresh='refresh'>
    </bulkNotarization>
  </div>
</div>`,
  data: () => ({
    filter: '',
    images: [],
    initialized: false,
  }),
  mounted: function () {
    this.poll();
  },
  components: {
    loading: LoadingComponent,
    notarizationTable: NotarizationTableComponent,
    bulkNotarization: BulkNotarizationComponent,
  },
  methods: {
    poll: function () {
      const self = this;
      setTimeout(function () {
        self.refresh(function () {
          self.initialized = true;
          self.poll();
        });
      }, ConfigurationService.PollInterval)
    },
    refresh: function (callback) {
      const self = this;
      NotarizationService.fetch(self.filter, function (images) {
        self.images = images;
        if (callback) {
          callback();
        }
      });
    },
    search: function (text) {
      this.filter = text;
      this.refresh();
    }
  },
};

export {ContainerComponent};
