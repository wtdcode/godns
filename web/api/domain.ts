import { get_api_server } from '@/api/env';

export interface Domain {
	domain_name: string;
	sub_domains: string[];
}

export async function get_domains(credentials: string): Promise<Domain[]> {
	if (credentials) {
		// make a GET request to the /api/auth endpoint via basic authentication
		const resp = await fetch(get_api_server() + '/api/v1/domains', {
			method: 'GET',
			headers: {
				'Authorization': `Basic ${credentials}`
			}
		})

		if (resp.status === 200) {
			return resp.json();
		}
	}

	return {} as Domain[];
}