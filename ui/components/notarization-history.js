import {NotarizationBadgeComponent} from "./notarization-badge.js";

const NotarizationHistoryComponent = {
  props: ['notarizations'],
  components: {
    notarizationBadge: NotarizationBadgeComponent,
  },
  template: `
<ul
  style='list-style-type: none; list-style-position: inside; padding-left: 0;'>
  <li
    v-for='notarization of notarizations'>
    <notarizationBadge v-bind:notarization='notarization'></notarizationBadge>
  </li>
</ul>`
};

export {NotarizationHistoryComponent};
