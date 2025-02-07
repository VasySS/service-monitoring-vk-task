import { BACKEND_URL } from '$lib';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	try {
		const containerData = await fetch(`${BACKEND_URL}/v1/statuses`).then((res) => res.json());
		return {
			containerData
		};
	} catch (err) {
		return {
			containerData: [],
			error: err
		};
	}
};
