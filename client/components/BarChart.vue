<script>
import { Bar, mixins } from 'vue-chartjs'
export default {
  mixins: [Bar, mixins.reactiveData],
  props: {
    dataset: {
      type : Object,
      requier: true
    }
  },
  extends: Bar,
  data() {
    return {
      chartData: {
        labels: [],
        datasets: [{
          label: 'Bar Dataset',
          backgroundColor: [
            'rgba(255, 60, 60, 0.3)',
            'rgba(60, 60, 255, 0.3)',
            'rgba(60, 255, 60, 0.3)',
            'rgba(160, 60, 60, 0.3)',
            'rgba(60, 160, 60, 0.3)',
            'rgba(60, 60, 160, 0.3)',
            'rgba(10, 60, 60, 0.3)',
            'rgba(60, 10, 60, 0.3)',
            'rgba(60, 60, 10, 0.3)',
          ],
          data: []
        }],
      },
      options: {
        scales: {
          yAxes: [{
            ticks: {
              beginAtZero: true,
            }
          }]
        },
        title: {
          display: true,
          text: '各スコア'
        },
      }
    }
  },
  mounted () {
    this.chartData.labels = Object.keys(this.dataset)
    this.chartData.datasets[0].data = Object.values(this.dataset)
    this.renderChart(this.chartData, this.options)
  },
  watch: {
    dataset: function() {
      this.chartData.labels = Object.keys(this.dataset)
      this.chartData.datasets[0].data = Object.values(this.dataset)
      this.renderChart(this.chartData, this.options)
    }
  }
}
</script>