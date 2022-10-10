import type { DateTimeFormat } from '@intlify/core-base'

export const zh_cn_message = {
  'video intelligence': '视频智能',
  sidebar: {
    home: '主页',
    dashboard: '仪表盘',
    organization: '组织',
    marketplace: '市场',
    camera: '摄像头管理',
    camera_rtsp: 'RTSP设备',
    camera_gb28181: 'GB28181设备',
    camera_group: '分组',
    camera_map: '地图',
    intelligence: '智能管理',
    databining: '数据联动',
    settings: '设置',
  },
  pages: {
    camera: {
      addcamera: '添加摄像头',
      tables: '表格',
      groups: '分组',
      map: '地图'
    }
  },
  camera: {
    profile: {
      serialNumber: '设备编号',
      name: '设备名称',
      status: '状态',
      description: '说明',
      created: '创建时间',
      info: '更多信息',
      channel: '码流',
      host: '网络地址',
      rtsp:{
        port: '端口号',
        username: '用户名',
        pwd: '密码',
      }
    },
    status: {
      online: '在线',
      offline: '离线'
    }
  },
  other: {
    search: '搜索',
    unknown: '未知'
  }
};

export const zh_cn_datetime:DateTimeFormat = {
  short: {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  },
  long: {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    weekday: 'long',
    hour: 'numeric',
    minute: 'numeric',
    hour12: true
  }
}
