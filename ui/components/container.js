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
	    v-bind:histories='histories'
	    v-on:refresh='refresh'
	    v-on:toggle-history='toggleHistory'>
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
    histories: {},
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
        if (JSON.stringify(self.images) !== JSON.stringify(images)) {
          self.images = images;
          self.histories = {};
        }
        if (callback) {
          callback();
        }
      });
    },
    search: function (text) {
      this.filter = text;
      this.refresh();
    },
    toggleHistory: function (image) {
      const self = this;
      if (this.histories[image.Image.Hash]) {
        this.histories = {};
      } else {
        NotarizationService.history(image.Image.Hash, function (history) {
          self.histories = {};
          self.histories[image.Image.Hash] = history;
        })
      }
    },
  },
};

export {ContainerComponent};
