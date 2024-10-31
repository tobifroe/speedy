<script lang="ts">
	import type { PageData } from './$types';
	import Chart from 'chart.js/auto';

	let { data }: { data: PageData } = $props();
	let canvas;
	setTimeout(() => {
		let c = new Chart(canvas, {
			type: 'line',
			options: {
				scales: {
					y: {
						title: {
							display: true,
							text: 'MB/s'
						}
					},
					x: {
						display: false
					}
				}
			},
			data: {
				labels: data.testResults.map((x) => new Date(x.CreatedAt).toISOString()),
				datasets: [
					{
						label: 'Speedtest Results in MB/s',
						data: data.testResults.map((x) => x.DownSpeed),
						fill: false
					}
				]
			}
		});
	}, 0);
</script>

<div>
	<canvas bind:this={canvas} id="myChart"></canvas>
</div>
