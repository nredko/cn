const SearchComponent = {
  template: `
<form class='d-inline'>
   <div class='input-group'>
      <input type='text' class='form-control' placeholder='search...' v-on:input='search'>
       <span class='input-group-append'></span>
    </div>
</form>`,
  methods: {
    search: function (input) {
      this.$emit('search', input.target.value);
    }
  }
};

export {SearchComponent};
