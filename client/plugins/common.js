import Vue from "vue";

Vue.mixin({
  methods: {
    GetURL: function() {
      const local = "http://localhost:8080/api/v1";
      return local;
    }
  }
});
