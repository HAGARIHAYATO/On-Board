<template>
  <div class="layer">
    <div class="assessModal">
      <p class="closeBtn" @click="closeMdl">×</p>
      <p class="description" v-if="evaluated">あなたの評価です。グラフにカーソルを合わせると更新します。</p>
      <p class="description" v-else>未評価です。グラフにカーソルを合わせると更新します。</p>
      <div class="chart">
        <Radar-Chart :dataset="radarData" :height="330" :width="660"/>
      </div>
      <div class="selectors">
        <p>
          機能の充実
          <select v-model="radarData['機能の充実度']">
            <option value=0>0</option>
            <option value=1>1</option>
            <option value=2>2</option>
            <option value=3>3</option>
            <option value=4>4</option>
            <option value=5>5</option>
            <option value=6>6</option>
            <option value=7>7</option>
            <option value=8>8</option>
            <option value=9>9</option>
            <option value=10>10</option>
          </select>
        </p>
        <p>
          UI/UX
          <select v-model="radarData['UI/UX']">
            <option value=0>0</option>
            <option value=1>1</option>
            <option value=2>2</option>
            <option value=3>3</option>
            <option value=4>4</option>
            <option value=5>5</option>
            <option value=6>6</option>
            <option value=7>7</option>
            <option value=8>8</option>
            <option value=9>9</option>
            <option value=10>10</option>
          </select>
        </p>
        <p>
          不都合な動作の少なさ
          <select v-model="radarData['不都合な動作の少なさ']">
            <option value=0>0</option>
            <option value=1>1</option>
            <option value=2>2</option>
            <option value=3>3</option>
            <option value=4>4</option>
            <option value=5>5</option>
            <option value=6>6</option>
            <option value=7>7</option>
            <option value=8>8</option>
            <option value=9>9</option>
            <option value=10>10</option>
          </select>
        </p>
        <p>
          内容の斬新さ
          <select v-model="radarData['内容の斬新さ']">
            <option value=0>0</option>
            <option value=1>1</option>
            <option value=2>2</option>
            <option value=3>3</option>
            <option value=4>4</option>
            <option value=5>5</option>
            <option value=6>6</option>
            <option value=7>7</option>
            <option value=8>8</option>
            <option value=9>9</option>
            <option value=10>10</option>
          </select>
        </p>
        <p>
          言語やFWのモダンさ
          <select v-model="radarData['言語やFWのモダンさ']">
            <option value=0>0</option>
            <option value=1>1</option>
            <option value=2>2</option>
            <option value=3>3</option>
            <option value=4>4</option>
            <option value=5>5</option>
            <option value=6>6</option>
            <option value=7>7</option>
            <option value=8>8</option>
            <option value=9>9</option>
            <option value=10>10</option>
          </select>
        </p>
      </div>
      <p class="postBtn" @click="postAssess()">評価を送信</p>
    </div>
  </div>
</template>
<script>
import RadarChart from "~/components/RadarChart.vue";
export default {
  data() {
    return {
      evaluated: true,
      APIURL: "",
      radarData: {
        "機能の充実度": 0,
        "UI/UX": 0,
        "不都合な動作の少なさ": 0,
        "内容の斬新さ": 0,
        "言語やFWのモダンさ": 0
      },
    }
  },
  components: {
    RadarChart
  },
  async mounted() {
    this.APIURL = this.GetURL();
    await this.$axios
      .get(this.APIURL + this.$route.path + "/login_user/assessment", {
        params: {
          user_id: this.$auth.user.ID
        }
      })
      .then(response => {
        console.log(response)
        this.assessmentInit(response.data.Assessment)
      })
      .catch(response => {
        console.error(response)
        this.evaluated = false
      })
  },
  methods: {
    assessmentInit: function(assess) {
      this.radarData["機能の充実度"] = assess.Function
      this.radarData["不都合な動作の少なさ"] = assess.BugSafe
      this.radarData["UI/UX"] = assess.UIX
      this.radarData["内容の斬新さ"] = assess.Content
      this.radarData["言語やFWのモダンさ"] = assess.MDN
    },
    postAssess: async function() {
      this.APIURL = this.GetURL();
      const data = new FormData();
      const headers = { "content-type": "multipart/form-data" };
      data.append("user_id", this.$auth.user.ID);
      data.append("function", this.radarData["機能の充実度"]);
      data.append("uix", this.radarData["UI/UX"]);
      data.append("bug_safe", this.radarData["不都合な動作の少なさ"]);
      data.append("content", this.radarData["内容の斬新さ"]);
      data.append("mdn", this.radarData["言語やFWのモダンさ"]);
      await this.$axios
        .post(this.APIURL + this.$route.path + "/assessment", data, {
          headers
        })
        .then(response => {
          this.$emit("loading")
        })
        .catch(response => console.error(response));
    },
    closeMdl: function() {
      this.$emit("close")
    }
  },
}
</script>
<style lang="scss" scoped>
.layer {
  position: fixed;
  top: 0;
  bottom: 0;
  background-color: rgba($color: white, $alpha: 0.8);
  width: 100%;
  height: 100%;
  z-index: 2;
}
.assessModal {
  width: 700px;
  height: 550px;
  margin: 0 auto;
  box-shadow: 0 0 2px black;
  background-color: white;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 3;
  border-radius: 5px;
  padding: 20px;
}
.closeBtn {
  position: absolute;
  top: 5px;
  right: 5px;
  border-radius: 50px;
  color: white;
  background-color: grey;
  height: 28px;
  width: 28px;
  cursor: pointer;
  font-weight: bold;
  font-size: 20px;
  line-height: 24px;
  text-align: center;
}
.chart {
  margin: 30px 0;
}
.selectors {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
}
.postBtn {
  margin: 20px auto;
  background-color: blue;
  border-radius: 5px;
  width: 200px;
  text-align: center;
  color: white;
  font-weight: bold;
  box-shadow: 0 2px 2px grey;
  cursor: pointer;
  &:active {
    box-shadow: 0 0 0 grey;
  }
}
.description {
  color: blue;
  font-size: 12px;
  margin-bottom: 10px;
}
</style>