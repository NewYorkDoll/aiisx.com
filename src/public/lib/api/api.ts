import { GetEventsQuery } from "~~/.nuxt/gql-sdk";



export type GithubEventNode = NonNullable<NonNullable<NonNullable<GetEventsQuery["githubevents"]["edges"]>[number]>["node"]>