<script lang="ts">
	import type { PageData } from './$types';

	interface Props {
		data: PageData;
	}
	let { data }: Props = $props();

	let containerData = data.containerData;
	let error = data.error;
</script>

<div class="flex min-h-screen flex-col items-center justify-center p-8">
	<h1 class="pb-10">Container pinger</h1>

	{#if error}
		<div class="mb-4 text-red-500">{error}</div>
	{:else if containerData.length === 0}
		<div>No containers found</div>
	{:else}
		<div class="w-full max-w-4xl overflow-x-auto rounded-lg shadow-lg">
			<table class="min-w-full border-collapse bg-white">
				<thead class="bg-gray-800 text-white">
					<tr>
						<th>ID</th>
						<th>IP</th>
						<th>Status</th>
						<th>Created at</th>
					</tr>
				</thead>

				<tbody class="divide-y divide-gray-400">
					{#each containerData as container}
						<tr class="transition-colors hover:bg-gray-500">
							<td>{container.containerID.slice(0, 12)}</td>
							<td>{container.ip}</td>
							<td>{container.status}</td>
							<td>{container.createdAt}</td>
							<!-- <td>
							{#if container.state === 'running'}
								<span style="color: green;">●</span>
							{:else}
								<span style="color: red;">●</span>
							{/if}
							{container.state}
						</td> -->
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	{/if}
</div>
