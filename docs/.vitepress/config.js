import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Rapide',
  description: 'Bullet Journal-style rapid logging for the terminal',
  base: '/rapide/',

  head: [
    ['meta', { name: 'theme-color', content: '#04B575' }],
  ],

  themeConfig: {
    logo: '/logo.svg',

    nav: [
      { text: 'Guide', link: '/getting-started' },
      { text: 'Reference', link: '/reference/commands' },
      { text: 'Roadmap', link: '/roadmap' },
      { text: 'GitHub', link: 'https://github.com/codevalve/rapide' },
    ],

    sidebar: [
      {
        text: 'Introduction',
        items: [
          { text: 'Getting Started', link: '/getting-started' },
        ],
      },
      {
        text: 'Guide',
        items: [
          { text: 'Syntax & Bullets', link: '/guide/syntax' },
          { text: 'TUI Interface', link: '/guide/tui' },
        ],
      },
      {
        text: 'Reference',
        items: [
          { text: 'CLI Commands', link: '/reference/commands' },
          { text: 'Configuration', link: '/reference/configuration' },
          { text: 'Git Sync', link: '/reference/git-sync' },
        ],
      },
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/codevalve/rapide' },
    ],

    editLink: {
      pattern: 'https://github.com/codevalve/rapide/edit/main/docs/:path',
      text: 'Edit this page on GitHub',
    },

    footer: {
      message: 'Built with the Charmbracelet toolchain 🗿',
      copyright: '© 2026 CodeValve',
    },

    search: {
      provider: 'local',
    },
  },
})
