import { zh_cn_message, zh_cn_datetime } from './zh-CN';

export type MessageSchema = typeof zh_cn_message
export type DataTimeSchema = typeof zh_cn_datetime

export const messages = {
  'zh-CN': zh_cn_message,
  'en-US': zh_cn_message
}

export const datetimeFormats = {
  'zh-CN': zh_cn_datetime,
  'en-US': zh_cn_datetime
};
