<template>
  <div class="app-generate-form">
    <div class="form-container jumbotron">
      <h1 class="display-5">Welcome to Elf!</h1>

      <form class="mt-4" @submit.prevent="generateUrl">
        <div class="form-group">
          <label for="exampleInputEmail1"
            >Enter a URL to generate its long URL</label
          >
          <input
            class="form-control"
            v-model="formData.originalUrl"
            aria-describedby="emailHelp"
            placeholder="Eg: https://www.facebook.com/"
          />
        </div>
        <div class="form-group">
          <label for="exampleInputEmail1">Expires After</label>
          <div class="input-group">
            <div class="input-group-prepend">
              <span
                class="input-group-text"
                id="validationTooltipUsernamePrepend"
                >Seconds</span
              >
            </div>
            <input
              type="number"
              class="form-control"
              v-model="formData.expiresAfter"
              aria-describedby="emailHelp"
              placeholder="Eg: 50 (Default: 60)"
              min="0"
              max="300"
            />
          </div>
        </div>
        <div class="form-group">
          <button class="btn btn-primary" href="#" role="button">
            Generate
          </button>
        </div>
      </form>

      <hr class="my-4" />
      <p class="lead" v-if="shortUrlRecieved">
        Generated Short URL:
        <a
          target="_"
          :href="`http://localhost:5000/${responseData.shortUrl}`"
          >{{ `http://localhost:5000/${responseData.shortUrl}` }}</a
        >
      </p>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
export default {
  name: "ElfGenerateForm",
  components: {},
  data() {
    return {
      shortUrlRecieved: false,
      formData: {
        originalUrl: null,
        expiresAfter: null
      },
      responseData: {
        originalUrl: "awxwaxwddawdaa",
        shortUrl: "70848b",
        createdAt: "2020-04-02T18:15:25.624026969+05:30",
        expiresAfter: "2020-04-02T18:16:25.624029244+05:30",
        hasExpired: false
      }
    };
  },
  methods: {
    generateUrl() {
      this.shortUrlRecieved = false;

      if (
        this.formData.originalUrl == null ||
        this.formData.originalUrl == ""
      ) {
        console.log("Please enter the URL before submitting the form.");
      } else {
        let finalFormData = {
          originalUrl: this.formData.originalUrl,
          expiresAfter: this.formData.expiresAfter || 60
        };

        console.log(finalFormData);

        this.axios
          .post("/api/generate", finalFormData)
          .then(res => {
            this.shortUrlRecieved = true;
            this.responseData = res.data.Payload;
          })
          .catch(err => {
            this.shortUrlRecieved = false;
            console.error(err);
          });
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.app-generate-form {
  padding-top: 5rem;
  display: flex;
  flex-direction: column;
  align-items: center;

  .form-container {
    min-width: 50%;
    background-color: #f8f8f8;
    border-radius: 1rem;
    box-shadow: 0px 0px 20px #eeeeee;
  }
}
</style>
