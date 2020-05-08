<script>
import { Pie, mixins } from 'vue-chartjs'
export default {
  mixins: [Pie, mixins.reactiveData],
  props: {
    dataset: {
      type : Object,
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
        title: {
          display: true,
          text: 'プロジェクト内言語別使用率'
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