<script>
import { Radar, mixins } from 'vue-chartjs'
export default {
  mixins: [Radar, mixins.reactiveData],
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
  extends: Radar,
  data() {
    return {
      dataframe: {},
      colors: [    
        'rgba(0, 0, 255, 0.4)',
      ],
      chartData: {
        labels: [],
        datasets: [{
          backgroundColor: 'rgba(0, 0, 255, 0.4)',
          borderWidth: 5,
          borderColor: 'rgba(0, 0, 255, 0.8)',
          data: [],
          pointBackgroundColor: 'rgba(0, 0, 255, 0.8)',
        }],
      },
      options: {
        animation: {
          animateScale: true,
        },
        scale: {
          ticks: {
            maxTicksLimit: 10,
            beginAtZero: true,
            min: 0,
            stepSize: 1,
          },
        },
        responsive: true,
        legend: {
          display: false,
        },
        title: {
          display: false
        },
        spanGaps: true,
      }
    }
  },
  mounted () {
    this.dataframe = this.dataset
    this.chartData.datasets[0].backgroundColor, this.chartData.datasets[0].borderColor = this.createColorList()
    this.chartData.labels = Object.keys(this.dataset)
    this.chartData.datasets[0].data = Object.values(this.dataframe)
    this.renderChart(this.chartData, this.options)
    this.$data._chart.canvas.onmouseenter = this.getMove
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
      return x
    }
  },
}
</script>