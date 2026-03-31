<template>
  <div class="markdown-body" v-html="renderedHtml"></div>
</template>

<script setup>
import { computed } from 'vue'
import MarkdownIt from 'markdown-it'
import hljs from 'highlight.js'
import 'highlight.js/styles/github-dark.css'

const props = defineProps({
  content: {
    type: String,
    default: ''
  }
})

const md = new MarkdownIt({
  html: true,
  linkify: true,
  typographer: true,
  highlight(str, lang) {
    if (lang && hljs.getLanguage(lang)) {
      try {
        return `<pre class="hljs"><code>${hljs.highlight(str, { language: lang, ignoreIllegals: true }).value}</code></pre>`
      } catch (_) {
        // ignore
      }
    }
    return `<pre class="hljs"><code>${md.utils.escapeHtml(str)}</code></pre>`
  }
})

const renderedHtml = computed(() => {
  if (!props.content) return ''
  return md.render(props.content)
})
</script>
