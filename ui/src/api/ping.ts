import { Get } from '@/utils/api'

const v1_url = 'api/v1'


export interface PingResData {
  ping: string;
}

export function Ping() {
  return Get<unknown, PingResData>(v1_url + '/ping');
}
