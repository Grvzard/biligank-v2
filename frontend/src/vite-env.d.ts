/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_BACKEND_URL: string
  readonly VITE_ADS_BACKEND: string
  readonly VITE_DONATE_LINK: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
