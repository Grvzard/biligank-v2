/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_BACKEND_URL: string
  readonly VITE_ADS_BACKEND: string
  // readonly VITE_DONATE_LINK: string
  readonly VITE_LIVEPUSH_BACKEND: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
