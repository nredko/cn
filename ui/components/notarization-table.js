import {NotarizationBadgeComponent} from "./notarization-badge.js";
import {NotarizationButtonComponent} from "./notarization-button.js";

const NotarizationTableComponent = {
  template: `
<table 
  class='table'>
  <thead 
    class='borderless' 
    style='border-top: hidden;'>
  <tr>
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
  		<notarizationBadge 
  		  v-bind:image='image'>
      </notarizationBadge>
  	</td>
  	<td>
      <notarizationButton 
        v-bind:image='image' 
        v-on='$listeners'>
       </notarizationButton>
  	</td>
  </tr>
  </tbody>
</table>`,
  components: {
    notarizationBadge: NotarizationBadgeComponent,
    notarizationButton: NotarizationButtonComponent
  },
  props: ['images'],
  data: () => ({
    images: [],
  }),
};

export {NotarizationTableComponent};
