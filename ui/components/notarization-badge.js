const NotarizationBadgeComponent = {
  template: `
<h4>
  <span
    v-if='image.Notarization.Status === "Notarized"'
    class='badge badge-pill badge-success'>
    Notarized
  </span>
  <span
    v-if='image.Notarization.Status === "Untrusted"'
    class='badge badge-pill badge-danger'>
    Un-trusted
  </span>
  <span
    v-if='image.Notarization.Status === "Unknown"'
    class='badge badge-pill badge-warning'>
    Unknown
  </span>
</h4>`,
  props: ['image'],
  data: function () {
    return {
      image: {},
    }
  }
};

export {NotarizationBadgeComponent};
