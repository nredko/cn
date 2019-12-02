import {SearchComponent} from "./search.js";

const NavbarComponent = {
  components: {
    search: SearchComponent,
  },
  template: `
<nav 
  class='navbar navbar-light'>
	<a 
	  href='#'
	  class='navbar-brand logo'
	  style='margin-left: 50px; vertical-align: middle'>
		<img 
		  src='assets/cn-color.eeadbabe.svg'
			width='200'
			height='60'
			class='d-inline-block align-top'
			alt=''>
	</a>
	<search v-on='$listeners'></search>
</nav>
`,
};

export {NavbarComponent}
