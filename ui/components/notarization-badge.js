const NotarizationBadgeComponent = {
  template: `
<h4>
  <span
    v-if='notarization.Status === "Notarized"'
    class='badge badge-pill badge-success'>
    Notarized
  </span>
  <span
    v-if='notarization.Status === "Untrusted"'
    class='badge badge-pill badge-danger'>
    Un-trusted
  </span>
  <span
    v-if='notarization.Status === "Unknown"'
    class='badge badge-pill badge-warning'>
    Unknown
  </span>
</h4>`,
  props: ['notarization'],
  data: function () {
    return {
      notarization: {},
    }
  }
};

export {NotarizationBadgeComponent};
