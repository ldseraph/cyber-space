import { Get } from '@/utils/api'

const v1_url = '/v1'


export interface PingResData {
  ping: string;
}

export function ping() {
  return Get<unknown, PingResData>(v1_url + '/ping');
}
