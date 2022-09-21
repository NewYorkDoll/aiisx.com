export interface BaseQuery {
  githubUser: {
    htmlurl: string;
    name: string;
    dio: string;
  };
}
export interface History {
  title: string;
  path: string;
  timestamp: string;
}

export interface State {
  base: BaseQuery | null;
  history: History[];
  sidebarCollapsed: boolean;
}
export const useBaseState = defineStore("state", {
  state: (): State => ({
    base: {
      githubUser: {
        htmlurl: "https://github.com/NewYorkDoll",
        name: "aiisx",
        dio: "NewYorkDoll",
      },
    },
    history: [],
    sidebarCollapsed: false,
  }),
  actions: {
    addToHistory(item: History) {
      // Truncate to a max size.
      if (this.history.length > 4) {
        this.history.shift();
      }

      // Remove any previous duplicates with the exact same path.
      for (let i = this.history.length - 1; i >= 0; i--) {
        if (this.history[i].path === item.path) {
          this.history.splice(i, 1);
        }
      }

      this.history.push(item);
    },
  },
});
