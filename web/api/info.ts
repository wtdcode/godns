import { get_api_server } from '@/api/env';

export interface Info {
	version: string;
	start_time: number;
	domains: number;
	sub_domains: number;
}

export async function get_info(credentials: string): Promise<Info> {
	// make a GET request to the /api/auth endpoint via basic authentication
	const resp = await fetch(get_api_server() + '/api/v1/info', {
		method: 'GET',
		headers: {
			'Authorization': `Basic ${credentials}`
		}
	})

	if (resp.status === 200) {
		return resp.json();
	}

	return {} as Info;
}

export function get_hours(timestamp: number): string {
	// compute the number of hours between the current time and the timestamp
	const current_time = Date.now() / 1000;
	const diff = (current_time - timestamp);
	const hours = (diff / 3600).toFixed(1);
	return `${hours} Hours`;
}

export function get_date(timestamp: number): string {
	// convert the timestamp to a human-readable date
	const date = new Date(timestamp * 1000);
	// convert date to YYYY-MM-DD HH:mm:ss format with the time in local timezone
	return date.toLocaleString();
}