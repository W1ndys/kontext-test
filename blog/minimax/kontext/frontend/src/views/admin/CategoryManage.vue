<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { getCategoryList, createCategory, updateCategory, deleteCategory } from '@/api/category'
import { formatDateTime } from '@/utils/time'
import type { Category } from '@/types'

const categories = ref<Category[]>([])
const loading = ref(true)
const showModal = ref(false)
const editingId = ref<number | null>(null)
const form = ref({ name: '', slug: '', sort: 0 })

onMounted(async () => {
  await fetchCategories()
})

async function fetchCategories() {
  try {
    categories.value = await getCategoryList()
  } catch (e) {
    console.error('Failed to load categories:', e)
  } finally {
    loading.value = false
  }
}

function openCreate() {
  editingId.value = null
  form.value = { name: '', slug: '', sort: 0 }
  showModal.value = true
}

function openEdit(cat: Category) {
  editingId.value = cat.id
  form.value = { name: cat.name, slug: cat.slug, sort: cat.sort }
  showModal.value = true
}

async function handleSubmit() {
  if (!form.value.name.trim()) {
    alert('请输入分类名称')
    return
  }

  try {
    if (editingId.value) {
      await updateCategory(editingId.value, form.value)
    } else {
      await createCategory(form.value)
    }
    showModal.value = false
    await fetchCategories()
  } catch (e) {
    alert('保存失败')
  }
}

async function handleDelete(id: number) {
  if (!confirm('确定要删除这个分类吗？')) return
  try {
    await deleteCategory(id)
    await fetchCategories()
  } catch (e) {
    alert('删除失败')
  }
}
</script>

<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold">分类管理</h2>
      <button @click="openCreate" class="btn btn-primary">新建分类</button>
    </div>

    <div class="bg-white rounded-xl shadow-md overflow-hidden">
      <table class="w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">名称</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">别名</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">排序</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="cat in categories" :key="cat.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">{{ cat.name }}</td>
            <td class="px-6 py-4 text-gray-500">{{ cat.slug }}</td>
            <td class="px-6 py-4 text-gray-500">{{ cat.sort }}</td>
            <td class="px-6 py-4 text-right">
              <button @click="openEdit(cat)" class="text-primary-600 hover:underline mr-4">编辑</button>
              <button @click="handleDelete(cat.id)" class="text-red-600 hover:underline">删除</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div v-if="loading" class="text-center py-12 text-gray-500">加载中...</div>
      <div v-if="!loading && categories.length === 0" class="text-center py-12 text-gray-500">
        暂无分类
      </div>
    </div>

    <div v-if="showModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl p-6 w-full max-w-md">
        <h3 class="text-xl font-bold mb-4">{{ editingId ? '编辑分类' : '新建分类' }}</h3>
        <div class="mb-4">
          <label class="block text-gray-700 mb-2">名称</label>
          <input v-model="form.name" type="text" class="input" placeholder="分类名称" />
        </div>
        <div class="mb-4">
          <label class="block text-gray-700 mb-2">别名</label>
          <input v-model="form.slug" type="text" class="input" placeholder="slug" />
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 mb-2">排序</label>
          <input v-model.number="form.sort" type="number" class="input" />
        </div>
        <div class="flex justify-end gap-3">
          <button @click="showModal = false" class="btn btn-secondary">取消</button>
          <button @click="handleSubmit" class="btn btn-primary">保存</button>
        </div>
      </div>
    </div>
  </div>
</template>
