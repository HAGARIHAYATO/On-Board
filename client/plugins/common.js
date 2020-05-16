import Vue from "vue";

import html2canvas from "html2canvas";
import pdfMake from "pdfmake/build/pdfmake";
Vue.component('html2canvas', html2canvas)
Vue.component('pdfmake', pdfMake)


const RATE = 2.83464566929;
const PAGE_WIDTH = 297 * RATE;
const PAGE_HEIGHT = 419 * RATE;
const CONTENT_WIDTH = 297 * RATE;
const CONTENT_HEIGHT = 419 * RATE;
const PAGE_MARGINS = [0 * RATE, 0 * RATE];

Vue.mixin({
  methods: {
    createPdfFromHtml: async function(element) {
      const pdfProps = await this.createPdfProps(element);
      this.createPdf(pdfProps);
    },
    createPdfProps: async function(element) {
      const options = {
        scale: 2
      };
      const canvas = await html2canvas(element, options);
    
      const dataUrl = canvas.toDataURL();
    
      const pdfProps = {
        dataUrl,
        pageSize: {
          width: PAGE_WIDTH,
          height: PAGE_HEIGHT
        },
        pageOrientation: "PORTRAIT",
        contentSize: {
          width: CONTENT_WIDTH,
          height: CONTENT_HEIGHT
        },
        pageMargins: PAGE_MARGINS
      };
    
      return pdfProps;
    },
    createPdf: function(pdfProps) {
      const { dataUrl, contentSize, pageMargins } = pdfProps;
      const pageSize = pdfProps.pageSize;
      const pageOrientation = pdfProps.pageOrientation;
    
      const documentDefinitions = {
        pageSize,
        pageOrientation,
        content: {
          image: dataUrl,
          ...contentSize
        },
        pageMargins
      };
    
      pdfMake.createPdf(documentDefinitions).download();
    },
    calcSize: function() {
      if (window.innerWidth > 400) {
        return 400
      } else {
        return window.innerWidth * 0.9
      }
    },
    GetURL: function() {
      // const local = "/api/v1";
      const local = "https://api.on-board-project.com/api/v1";
      return local;
    },
    AlertMessage: function() {
      const message = "画像サイズが大きすぎます。(最大3MB)";
      return message;
    },
    FetchGitInfo: function(author, token) {
      if (token === "") return ""
      const github = "https://api.github.com"
      const user = "/users/"
      const orgs = "/orgs/"
      let endpoint = ""
      switch (author) {
        case 0:
        endpoint = github + user + token + "/repos"
        break;
        case 1:
        endpoint = github + orgs + token + "/repos"
        break;
      }
      return endpoint
    },
    RewriteTime: function(time) {
      if (time) {
        let fixedTime = ""
        fixedTime = time.split("T")
        fixedTime = fixedTime[0].split("-")
        fixedTime = fixedTime[0] + "年" + fixedTime[1] + "月" + fixedTime[2] + "日"
        return fixedTime
      }
    }
  }
});