import { defineConfig } from "vitepress";

export default defineConfig({
  title: "PhenoDevOps",
  description: "Phenotype DevOps automation, deployment, and operations workspace.",
  base: process.env.GITHUB_PAGES === "true" ? "/PhenoDevOps/" : "/",
  cleanUrls: true,
  themeConfig: {
    logo: { text: "PhenoDevOps" },
    nav: [
      { text: "Overview", link: "/" },
      { text: "Guide", link: "/guide" },
      { text: "Pipelines", link: "/pipelines" },
      { text: "Operations", link: "/operations" },
      { text: "GitHub", link: "https://github.com/KooshaPari/PhenoDevOps" },
    ],
    sidebar: [
      {
        text: "PhenoDevOps",
        items: [
          { text: "Overview", link: "/" },
          { text: "Guide", link: "/guide" },
          { text: "Pipelines", link: "/pipelines" },
          { text: "Operations", link: "/operations" },
        ],
      },
    ],
    socialLinks: [{ icon: "github", link: "https://github.com/KooshaPari/PhenoDevOps" }],
    search: {
      provider: "local",
    },
  },
});
