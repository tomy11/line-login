<template>
  <div>
    <div class="text-center">
      <h3>{{ name }}</h3>
      <h5>{{ userID }}</h5>
      <img :src="pictureUrl" with="100" height="100" class="rounded" alt="" />
        <br>
        <br>
      <div class="text-center">
        <button
          type="submit"
          class="btn btn-light"
          @click="getAccessToken"
          :disabled="loggedIn === false"
        >
          AccessToken
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import liff from "@line/liff";
export default {
  name: "Profile",
  components: {},
  data: () => ({
    loggedIn: false,
    inClient: false,
    name: "",
    pictureUrl: "",
    accessToken: "",
    os: "",
    language: "",
    liff_version: "",
    userID: "",
  }),
  created: function () {
    this.loggedIn = liff.isLoggedIn();
    this.inClient = liff.isInClient();
    liff
      .getProfile()
      .then((profile) => {
        this.name = profile.displayName;
        this.pictureUrl = profile.pictureUrl;
        this.userID = profile.userId;
      })
      .catch((err) => {
        console.log(`Error at getProfile: ${err}`);
        this.name = "";
      });
  },
  methods:{
    getAccessToken: function() {
        this.accessToken = liff.getAccessToken()
        const text = `Your Access Token is [${this.accessToken}]`
        console.log(text)
    },
  }
};
</script>