<template>
<div id="header"></div>
<div class="container-fluid">
    <section class="h-100">
      <div class="container h-100">
        <div class="row justify-content-sm-center h-100">
          <div class="col-xxl-4 col-xl-5 col-lg-5 col-md-7 col-sm-9">
            <div class="card shadow-lg">
              <div class="card-body p-5">
                <h1 class="fs-4 card-title fw-bold mb-4">Login</h1>
                <Login v-if="loggedIn === false" />
                <Profile v-else />
                <br><br>
                <div class="text-center">
                  <button type="submit" class="btn btn-light" @click="doLogout" v-show="loggedIn === true">logout</button>
                </div>
              </div>
              <div class="card-footer py-3 border-0">
                <div class="text-center">
                  account? <a href="#" class="text-dark">register</a>
                </div>
              </div>
            </div>
            <div class="text-center mt-5 text-muted">
     
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
  
</template>
<script>
import "bootstrap/dist/css/bootstrap.min.css";
import liff from '@line/liff';
import Login from './components/Login.vue'
import Profile from './components/Profile.vue'
export default {
  name: 'App',
  components: {
    Login,
    Profile
  },
  data: () => ({
    loggedIn: false,
    initialized: false,
  }),
  created: function() {
  },
  mounted: async function() {
   await this.initializeLiff()
  },
  methods: {
     initializeLiff: async function () {
      await liff.init({liffId: process.env.VUE_APP_LIFF_ID})
      .then(() => {
        console.log('init app')
        this.initialized = true
        if(liff.isLoggedIn()){
          console.log('is login')
          this.loggedIn = true
        }else{
          console.log('log out')
        }
      }).catch(err=>{throw err});
    },
    doLogout: function() {
      if (liff.isLoggedIn()) {
        liff.logout()
        this.loggedIn = false
        window.location.reload()
      }
    }
  }
};
</script>
<style>
#header{
  margin-top: 50px;
}
.btn-primary{
    background-color: #00B900 !important;
}
</style>
