<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { getTagList, createTag, updateTag, deleteTag } from '@/api/tag'
import type { Tag } from '@/types'

const tags = ref<Tag[]>([])
const loading = ref(true)
const showModal = ref(false)
const editingId = ref<number | null>(null)
const form = ref({ name: '', slug: '' })

onMounted(async () => {
  await fetchTags()
})

async function fetchTags() {
  try {
    tags.value = await getTagList()
  } catch (e) {
    console.error('Failed to load tags:', e)
  } finally {
    loading.value = false
  }
}

function openCreate() {
  editingId.value = null
  form.value = { name: '', slug: '' }
  showModal.value = true
}

function openEdit(tag: Tag) {
  editingId.value = tag.id
  form.value = { name: tag.name, slug: tag.slug }
  showModal.value = true
}

async function handleSubmit() {
  if (!form.value.name.trim()) {
    alert('请输入标签名称')
    return
  }

  try {
    if (editingId.value) {
      await updateTag(editingId.value, form.value)
    } else {
      await createTag(form.value)
    }
    showModal.value = false
    await fetchTags()
  } catch (e) {
    alert('保存失败')
  }
}

async function handleDelete(id: number) {
  if (!confirm('确定要删除这个标签吗？')) return
  try {
    await deleteTag(id)
    await fetchTags()
  } catch (e) {
    alert('删除失败')
  }
}
</script>

<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold">标签管理</h2>
      <button @click="openCreate" class="btn btn-primary">新建标签</button>
    </div>

    <div class="bg-white rounded-xl shadow-md overflow-hidden">
      <table class="w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">名称</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">别名</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="tag in tags" :key="tag.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">{{ tag.name }}</td>
            <td class="px-6 py-4 text-gray-500">{{ tag.slug }}</td>
            <td class="px-6 py-4 text-right">
              <button @click="openEdit(tag)" class="text-primary-600 hover:underline mr-4">编辑</button>
              <button @click="handleDelete(tag.id)" class="text-red-600 hover:underline">删除</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div v-if="loading" class="text-center py-12 text-gray-500">加载中...</div>
      <div v-if="!loading && tags.length === 0" class="text-center py-12 text-gray-500">
        暂无标签
      </div>
    </div>

    <div v-if="showModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl p-6 w-full max-w-md">
        <h3 class="text-xl font-bold mb-4">{{ editingId ? '编辑标签' : '新建标签' }}</h3>
        <div class="mb-4">
          <label class="block text-gray-700 mb-2">名称</label>
          <input v-model="form.name" type="text" class="input" placeholder="标签名称" />
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 mb-2">别名</label>
          <input v-model="form.slug" type="text" class="input" placeholder="slug" />
        </div>
        <div class="flex justify-end gap-3">
          <button @click="showModal = false" class="btn btn-secondary">取消</button>
          <button @click="handleSubmit" class="btn btn-primary">保存</button>
        </div>
      </div>
    </div>
  </div>
</template>
