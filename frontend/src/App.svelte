<script>
    import {Router, Route} from "svelte-routing";


    import Navbar from './components/Navbar.svelte';
    import About from "./components/About.svelte";
    import Article from "./components/Article.svelte";
    import Home from "./components/Home.svelte";
    import Works from "./components/Works.svelte";
    import Error from "./components/Error.svelte";

    // Used for SSR. A falsy value is ignored by the Router.
    export let url = "";
</script>


<Router url="{url}">
    <Navbar>
    </Navbar>

    <div>
        <Route path="/" component="{Home}"/>

        <Route path="/blog" component="{Article}"/>

        <Route path="/about" component="{About}"/>

        <Route path="/works" component="{Works}"/>

        <Route path="**" component="{Error}"/>
    </div>


</Router>
