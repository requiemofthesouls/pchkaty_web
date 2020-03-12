<script>
    import {Link} from "svelte-routing";
</script>

<nav class="uk-navbar-container uk-navbar-transparent" uk-navbar>
    <div class="uk-navbar-left uk-margin-left">
        <Link to="/">
            <img data-src="./images/cat.png" width="50" height="50" alt="" uk-img>
        </Link>
    </div>


    <div class="uk-navbar-right uk-margin-right">
        <ul class="uk-navbar-nav uk-text-bold">
            <li>
                <Link to="/works">Works</Link>
            </li>
            <li>
                <Link to="/blog">Thoughts</Link>
            </li>
            <li>
                <Link to="/about">About</Link>
            </li>
        </ul>
    </div>
</nav>


<style>
    .uk-navbar-container {
        background-color: var(--primary-color);
    }
</style>
