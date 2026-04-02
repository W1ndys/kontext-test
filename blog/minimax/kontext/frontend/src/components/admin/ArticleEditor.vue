<script setup lang="ts">
import { ref, watch } from 'vue'
import { handlePaste } from '@/utils/clipboard'

const props = defineProps<{
  modelValue: string
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
  (e: 'paste', url: string): void
}>()

const textareaRef = ref<HTMLTextAreaElement | null>(null)

const value = ref(props.modelValue)

watch(() => props.modelValue, (newVal) => {
  value.value = newVal
})

watch(value, (newVal) => {
  emit('update:modelValue', newVal)
})

function onPaste(event: ClipboardEvent) {
  handlePaste(event, (url) => {
    emit('paste', url)
  })
}

function insertText(before: string, after: string = '') {
  const textarea = textareaRef.value
  if (!textarea) return

  const start = textarea.selectionStart
  const end = textarea.selectionEnd
  const selected = value.value.substring(start, end)

  const newText = before + selected + after
  value.value = value.value.substring(0, start) + newText + value.value.substring(end)

  setTimeout(() => {
    textarea.focus()
    textarea.setSelectionRange(start + before.length, start + before.length + selected.length)
  }, 0)
}
</script>

<template>
  <div class="border border-gray-300 rounded-lg overflow-hidden">
    <div class="bg-gray-50 px-3 py-2 border-b border-gray-300 flex gap-2">
      <button type="button" @click="insertText('**', '**')" class="px-2 py-1 text-sm bg-white border rounded hover:bg-gray-100" title="粗体">
        <strong>B</strong>
      </button>
      <button type="button" @click="insertText('*', '*')" class="px-2 py-1 text-sm bg-white border rounded hover:bg-gray-100" title="斜体">
        <em>I</em>
      </button>
      <button type="button" @click="insertText('# ')" class="px-2 py-1 text-sm bg-white border rounded hover:bg-gray-100" title="标题">
        H
      </button>
      <button type="button" @click="insertText('[', '](url)')" class="px-2 py-1 text-sm bg-white border rounded hover:bg-gray-100" title="链接">
        Link
      </button>
      <button type="button" @click="insertText('`', '`')" class="px-2 py-1 text-sm bg-white border rounded hover:bg-gray-100" title="代码">
        Code
      </button>
      <button type="button" @click="insertText('\n```\n', '\n```\n')" class="px-2 py-1 text-sm bg-white border rounded hover:bg-gray-100" title="代码块">
        Block
      </button>
      <button type="button" @click="insertText('- ')" class="px-2 py-1 text-sm bg-white border rounded hover:bg-gray-100" title="列表">
        List
      </button>
      <button type="button" @click="insertText('\n> ')" class="px-2 py-1 text-sm bg-white border rounded hover:bg-gray-100" title="引用">
        Quote
      </button>
    </div>
    <textarea
      ref="textareaRef"
      v-model="value"
      @paste="onPaste"
      rows="20"
      class="w-full p-4 font-mono text-sm resize-none focus:outline-none"
      placeholder="支持 Markdown 语法，粘贴图片自动上传..."
    ></textarea>
  </div>
</template>
