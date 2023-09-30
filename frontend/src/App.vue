<template>
    <div>
        <h1>Hello world!</h1>
        <p>
            This is a template project which uses 
            <button @click="openLink('https://github.com/webview/webview')">Webview</button> for backend and 
            <button @click="openLink('https://github.com/vuejs/core')">Vue</button> for frontend.
        </p>
    </div>
</template>

<script>
export default {
    methods: {
        openLink(link) {
            openAnotherBrowser(link) // <= Bound to openLink() in main.go
        }
    }
}
</script>
