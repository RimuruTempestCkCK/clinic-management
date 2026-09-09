<script>
    import '../app.css';
    import Sidebar from '$lib/components/Sidebar.svelte';
    import Header from '$lib/components/Header.svelte';
    import Footer from '$lib/components/Footer.svelte';
    
    // We need to load the template JS scripts after components are mounted
    import { onMount } from 'svelte';
    
    let { children } = $props();
    
    onMount(() => {
        // Load runtime and vendor scripts
        const scripts = [
            '/runtime.js',
            '/vendor-fullcalendar.js',
            '/vendor-chartjs.js',
            '/vendors.js',
            '/2026.js'
        ];
        
        scripts.forEach(src => {
            const script = document.createElement('script');
            script.src = src;
            script.defer = true;
            document.body.appendChild(script);
        });
    });
</script>

<!-- The shell layout required by Adminator template -->
<div class="shell">
    <Sidebar />
    <div class="main">
        <Header />
        <main class="content">
            {@render children()}
        </main>
        <Footer />
    </div>
</div>
