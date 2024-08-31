export enum onlineStatusE { On, Off, Error, Unknown }

export interface StreamLog {
  title: string
  c_ts: number
  last_update: number
  watched_num: number
  area_name: string
  cover: string
}

export interface StreamerInfo {
  uid: number
  area_name: string
  parent_name: string
  roomid: number
  short_roomid: number
  uname: string
  first_log_ts: number
}
