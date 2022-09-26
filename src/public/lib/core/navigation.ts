import { NIcon } from "naive-ui";
import { Component, VNode } from "nuxt/dist/app/compat/capi";
import { RouterLink } from "vue-router";

function renderIcon(icon: Component | VNode | HTMLElement) {
  return () => h(NIcon, { style: "margin-top: -3px" }, { default: () => icon });
}

function renderLink(target: Record<string, any>, title: string, icon: any) {
  return {
    label: () =>
      h(
        RouterLink,
        { to: target },
        {
          default: () => title,
        }
      ),
    key: target.name,
    icon: renderIcon(icon),
  };
}

export const menuOptions = [
  { to: { name: "index" }, name: "Home", alias: "main" },
  { to: { name: "posts" }, name: "Posts", alias: "posts" },
  { to: { name: "repost" }, name: "repost", alias: "repost" },
];

export const branchMenuOptions = [
  { to: { name: "index" }, name: "master" },
  { to: { name: "util-gists" }, name: "feature/list-gists" },
];

export const adminSidebarOptions = [
  renderLink(
    { name: "admin" },
    "Admin Home",
    h(
      "div", // type
      { class: "i-mdi-home" }, // props
      []
    )
  ),
  renderLink(
    { name: "admin-posts" },
    "Blog Posts",
    h(
      "div", // type
      { class: "i-mdi-book-open-page-variant-outline" }, // props
      []
    )
  ),
  renderLink(
    { name: "admin-banner-builder" },
    "Banner Builder",
    h(
      "div", // type
      { class: "i-mdi-image-edit-outline" }, // props
      []
    )
  ),
  { key: "divider-1", type: "divider" },
  ...menuOptions.map((option) =>
    renderLink(
      option.to,
      option.name,
      h(
        "div", // type
        { class: "i-mdi-link" }, // props
        []
      )
    )
  ),
  { key: "divider-2", type: "divider" },
  {
    label: () =>
      h("a", { href: "/-/auth/logout" }, { default: () => "Logout" }),
    key: "logout",
    icon: renderIcon(
      h(
        "div", // type
        { class: "i-mdi-logout" }, // props
        []
      )
    ),
  },
];
