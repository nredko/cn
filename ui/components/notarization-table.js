import {NotarizationBadgeComponent} from "./notarization-badge.js";
import {NotarizationButtonComponent} from "./notarization-button.js";
import {NotarizationHistoryComponent} from "./notarization-history.js";

const NotarizationTableComponent = {
  template: `
<table
  class='table'>
  <thead
    class='borderless'
    style='border-top: hidden;'>
  <tr>
    <th></th>
	<th
  	  scope='col'>
  	  Docker Image
  	</th>
	<th
  	  scope='col'>
  	  Notarization Status
  	</th>
  </tr>
  </thead>
  <tbody
    id='notarization-table'>
  <tr
    v-for='image of images'>
  	<td>
  	  <h4>{{ image.Image.Name }}</h4>
  	</td>
  	<td>
	  <notarizationHistory
	    v-if='histories[image.Image.Hash]'
	    v-bind:notarizations='histories[image.Image.Hash]'>
      </notarizationHistory>
		<notarizationBadge
		  v-if='!histories[image.Image.Hash]'
		  v-bind:notarization='image.Notarization'>
      </notarizationBadge>
  	</td>
  	<td>
      <notarizationButton
        v-bind:image='image'
        v-on='$listeners'>
       </notarizationButton>
  	</td>
	 <td>
      <button
        role='button'
	    type='button'
	    class='btn btn-outline-primary'
        v-on:click='toggleHistory(image)'>
        Toggle History
      </button>
    </td>
  </tr>
  </tbody>
</table>`,
  components: {
    notarizationBadge: NotarizationBadgeComponent,
    notarizationButton: NotarizationButtonComponent,
    notarizationHistory: NotarizationHistoryComponent,
  },
  props: ['images', 'histories'],
  methods: {
    toggleHistory: function (image) {
      this.$emit("toggle-history", image);
    }
  }
};

export {NotarizationTableComponent};
