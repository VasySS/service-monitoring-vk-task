<script lang="ts">
	import type { PageData } from './$types';

	interface ContainerData {
		ip: string;
		status: string;
		createdAt: string;
		lastSuccess: string;
	}

	interface Props {
		data: PageData;
	}
	let { data }: Props = $props();

	let containerData = data.containerData as ContainerData[];
	let error = data.error;
</script>

<div class="flex min-h-screen flex-col items-center justify-center p-8">
	<h1 class="pb-10 text-2xl">Container pinger</h1>

	{#if error}
		<div class="mb-4 text-red-500">{error}</div>
	{:else if containerData.length === 0}
		<div>No containers found</div>
	{:else}
		<div class="w-full max-w-4xl overflow-x-auto rounded-lg shadow-lg">
			<table class="min-w-full border-collapse bg-white">
				<thead class="bg-gray-800 text-white">
					<tr>
						<th class="p-1.5">IP</th>
						<th>Status</th>
						<th>Created at</th>
						<th>Last success</th>
					</tr>
				</thead>

				<tbody class="divide-y divide-gray-400">
					{#each containerData as container}
						<tr class="text-center transition-colors hover:bg-gray-500">
							<td class="p-1.5">{container.ip}</td>
							<td>{container.status}</td>
							<td>{new Date(container.createdAt).toLocaleString()}</td>
							<td>{new Date(container.lastSuccess).toLocaleString()}</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	{/if}
</div>
