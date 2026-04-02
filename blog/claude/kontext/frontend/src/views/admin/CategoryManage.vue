<template>
  <AdminLayout>
    <div class="max-w-3xl">
      <h2 class="text-xl font-bold text-gray-800 mb-6">分类管理</h2>

      <!-- Add Form -->
      <div class="bg-white rounded-xl border border-gray-100 shadow-sm p-6 mb-6">
        <h3 class="text-sm font-semibold text-gray-700 mb-3">
          {{ editingCategory ? '编辑分类' : '添加分类' }}
        </h3>
        <form @submit.prevent="handleSubmit" class="flex gap-3">
          <input
            v-model="formName"
            type="text"
            placeholder="输入分类名称"
            required
            class="flex-1 px-4 py-2 border border-gray-300 rounded-lg text-gray-800 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
          />
          <button
            type="submit"
            :disabled="saving"
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-60"
          >
            {{ saving ? '保存中...' : (editingCategory ? '更新' : '添加') }}
          </button>
          <button
            v-if="editingCategory"
            type="button"
            @click="cancelEdit"
            class="px-4 py-2 bg-gray-100 text-gray-600 text-sm font-medium rounded-lg hover:bg-gray-200 transition-colors"
          >
            取消
          </button>
        </form>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="space-y-3">
        <div v-for="i in 5" :key="i" class="bg-white rounded-xl p-4 animate-pulse flex items-center gap-4">
          <div class="h-5 bg-gray-200 rounded flex-1"></div>
          <div class="h-5 bg-gray-200 rounded w-16"></div>
        </div>
      </div>

      <!-- Category List -->
      <div v-else class="bg-white rounded-xl border border-gray-100 shadow-sm overflow-hidden">
        <div v-if="categories.length === 0" class="p-8 text-center text-gray-400">
          暂无分类，请在上方添加
        </div>
        <div v-else class="divide-y divide-gray-100">
          <div
            v-for="cat in categories"
            :key="cat.id"
            class="flex items-center justify-between px-6 py-4 hover:bg-gray-50 transition-colors"
          >
            <div class="flex items-center gap-3">
              <span class="text-gray-800 font-medium">{{ cat.name }}</span>
              <span class="text-xs text-gray-400 bg-gray-100 px-2 py-0.5 rounded-full">
                {{ cat.article_count || 0 }} 篇文章
              </span>
            </div>
            <div class="flex items-center gap-2">
              <button
                @click="startEdit(cat)"
                class="text-sm text-blue-600 hover:text-blue-700 transition-colors"
              >
                编辑
              </button>
              <button
                @click="handleDelete(cat)"
                class="text-sm text-red-500 hover:text-red-600 transition-colors"
              >
                删除
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div
      v-if="showDeleteModal"
      class="fixed inset-0 bg-black/50 z-50 flex items-center justify-center px-4"
      @click.self="showDeleteModal = false"
    >
      <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-sm">
        <h3 class="text-lg font-semibold text-gray-800 mb-2">确认删除</h3>
        <p class="text-gray-500 text-sm mb-6">
          确定要删除分类「{{ deleteTarget?.name }}」吗？此操作不可撤销。
        </p>
        <div class="flex justify-end gap-3">
          <button
            @click="showDeleteModal = false"
            class="px-4 py-2 text-sm text-gray-600 bg-gray-100 rounded-lg hover:bg-gray-200 transition-colors"
          >
            取消
          </button>
          <button
            @click="confirmDelete"
            :disabled="deleting"
            class="px-4 py-2 text-sm text-white bg-red-500 rounded-lg hover:bg-red-600 transition-colors disabled:opacity-60"
          >
            {{ deleting ? '删除中...' : '确认删除' }}
          </button>
        </div>
      </div>
    </div>
  </AdminLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import AdminLayout from '@/components/AdminLayout.vue'
import { getCategories, createCategory, updateCategory, deleteCategory } from '@/api/category'

const categories = ref([])
const loading = ref(true)
const formName = ref('')
const editingCategory = ref(null)
const saving = ref(false)

const showDeleteModal = ref(false)
const deleteTarget = ref(null)
const deleting = ref(false)

async function fetchCategories() {
  loading.value = true
  try {
    const res = await getCategories()
    categories.value = res.data || res || []
  } catch (err) {
    console.error('Failed to fetch categories:', err)
  } finally {
    loading.value = false
  }
}

function startEdit(cat) {
  editingCategory.value = cat
  formName.value = cat.name
}

function cancelEdit() {
  editingCategory.value = null
  formName.value = ''
}

async function handleSubmit() {
  if (!formName.value.trim()) return
  saving.value = true
  try {
    if (editingCategory.value) {
      await updateCategory(editingCategory.value.id, { name: formName.value.trim() })
    } else {
      await createCategory({ name: formName.value.trim() })
    }
    formName.value = ''
    editingCategory.value = null
    fetchCategories()
  } catch (err) {
    console.error('Failed to save category:', err)
    alert('操作失败：' + (err.response?.data?.error || err.message))
  } finally {
    saving.value = false
  }
}

function handleDelete(cat) {
  deleteTarget.value = cat
  showDeleteModal.value = true
}

async function confirmDelete() {
  if (!deleteTarget.value) return
  deleting.value = true
  try {
    await deleteCategory(deleteTarget.value.id)
    showDeleteModal.value = false
    deleteTarget.value = null
    fetchCategories()
  } catch (err) {
    console.error('Failed to delete category:', err)
    alert('删除失败：' + (err.response?.data?.error || err.message))
  } finally {
    deleting.value = false
  }
}

onMounted(() => {
  fetchCategories()
})
</script>
