<script>
import { Pie, mixins } from 'vue-chartjs'
export default {
  mixins: [Pie, mixins.reactiveData],
  props: {
    dataset: {
      type : Object,
      requier: true
    },
    text: {
      type : String,
      requier: true
    }
  },
  extends: Pie,
  data() {
    return {
      chartData: {
        labels: [],
        datasets: [{
          backgroundColor: [
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
          ],
          data: []
        }],
      },
      options: {
        title: {
          display: true,
          text: this.text
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