import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, params }) => {
	const res = await fetch('http://localhost:8080/all');
	const testResults = await res.json();
	console.log(testResults.Results);
	return {
		testResults: testResults.Results
	};
	error(404, 'Not found');
};
