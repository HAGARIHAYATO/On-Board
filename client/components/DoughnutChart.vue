<script>
import { Doughnut, mixins } from 'vue-chartjs'
export default {
  mixins: [Doughnut, mixins.reactiveData],
  props: {
    dataset: {
      type : Object,
      requier: true
    },
    width: {
      type: Number,
      requier: true
    },
    height: {
      type: Number,
      requier: true
    },
    text: {
      type : String,
      requier: true
    }
  },
  extends: Doughnut,
  data() {
    return {
      dataframe: {},
      colors: [
        'rgba(255, 60, 60, 0.3)',
        'rgba(60, 60, 255, 0.3)',
        'rgba(60, 255, 60, 0.3)',
        'rgba(255, 255, 60, 0.3)',
        'rgba(60, 255, 255, 0.3)',
        'rgba(255, 60, 255, 0.3)',
        'rgba(255, 60, 60, 0.6)',
        'rgba(60, 60, 255, 0.6)',
        'rgba(60, 255, 60, 0.6)',
        'rgba(255, 255, 60, 0.6)',
        'rgba(60, 255, 255, 0.6)',
        'rgba(255, 60, 255, 0.6)',
        'rgba(255, 60, 60, 1)',
        'rgba(60, 60, 255, 1)',
        'rgba(60, 255, 60, 1)',
        'rgba(255, 255, 60, 1)',
        'rgba(60, 255, 255, 1)',
        'rgba(255, 60, 255, 1)',
        'rgba(255, 60, 60, 0.1)',
        'rgba(60, 60, 255, 0.1)',
        'rgba(60, 255, 60, 0.1)',
        'rgba(255, 255, 60, 0.1)',
        'rgba(60, 255, 255, 0.1)',
        'rgba(255, 60, 255, 0.1)',
      ],
      chartData: {
        labels: [],
        datasets: [{
          backgroundColor: [],
          borderWidth: 1,
          data: []
        }],
      },
      options: {
        cutoutPercentage: 80,
        animation: {
          animateScale: true,
        },
        responsive: true,
        legend: {
          position: 'bottom'
        },
        title: {
          display: true,
          text: this.text
        },
      }
    }
  },
  mounted () {
    this.dataframe = this.dataset
    this.chartData.datasets[0].backgroundColor = this.createColorList()
    this.chartData.labels = Object.keys(this.dataset)
    this.chartData.datasets[0].data = Object.values(this.dataframe)
    this.renderChart(this.chartData, this.options)
    this.$data._chart.canvas.onmouseenter = this.getMove
  },
  watch: {
    dataset: function() {
      this.dataframe = this.dataset
    }
  },
  watch: {
    dataframe: function() {
      this.dataframe = this.dataset
      this.chartData.labels = Object.keys(this.dataframe)
      this.chartData.datasets[0].data = Object.values(this.dataframe)
      this.renderChart(this.chartData, this.options)
    }
  },
  methods: {
    getMove: function(e) {
      this.chartData.datasets[0].backgroundColor = this.createColorList()
      const obj = this.dataframe
      this.dataframe = {}
      this.dataframe = obj
    },
    createColorList: function() {
      let x = this.colors
      x = x.map(function(a){return {weight:Math.random(), value:a}})
        .sort(function(a, b){return a.weight - b.weight})
        .map(function(a){return a.value});
      x.unshift("lightgrey");
      return x
    }
  },
}
</script>