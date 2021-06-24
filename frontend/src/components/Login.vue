<template>
  <div>

      <div class="mb-3">
        <p class="form-text text-muted mb-3">
          By registering you agree with our terms and condition.
        </p>
        <div class="align-items-center">
          <div class="d-grid gap-2">
            <button type="submit" class="btn btn-primary" @click="doLogin">line login</button>
          </div>
        </div>
      </div>
  
  </div>
</template>
<script>
import liff from "@line/liff";
export default {
  name: "Login",
  components: {},
  data: () => ({
    
  }),
  created: function () {this.loggedIn = true},
  mounted: function () {},
  methods: {
    doLogin: async function () {
      console.log('click login')
      if(liff.isLoggedIn()) {
        this.loggedIn = true
        let getProfile = await liff.getProfile();
        console.log(getProfile);
      }else{
        this.loggedIn = true
        liff.login();
      }
      
    },
  },
};
</script>