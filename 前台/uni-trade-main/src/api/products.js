import httpInstance from '@/utils/https'
import { useUserStore } from '@/store/userStore'

/**
 * 获取分类
 * @returns
 */
export const getCategoryAPI = () => {
  return httpInstance({
    url: '/home/category'
  })
}

/**
 * 获取商品列表
 * @param {商品的分类号码} category
 * @param {当前页数} page
 * @param {每页最大商品数 默认为12} limit
 * @returns
 */
export const getProductsListAPI = (category, page, limit, searchQuery) => {
  return httpInstance({
    url: '/products',
    params: {
      category,
      page,
      limit,
      searchQuery: searchQuery || undefined
    }
  })
}

//根据筛选条件筛选商品
// export const getFilteredProductsAPI = (data) => {
//   return httpInstance({
//     url: '/products',
//     method: 'POST',
//     data: data
//   })
// }

//根据筛选条件筛选商品
export function getFilteredProductsAPI({
  categoryID = 0,
  area = '',
  city = '',
  deliveryMethod = '',
  priceMax = 0,
  priceMin = 0,
  province = '',
  publishDate = [],
  shippingCost = 0,
  page = 1,
  limit = 12,
  searchQuery = ''
} = {}) {
  return httpInstance({
    url: '/product/select',
    params: {
      categoryID: categoryID || undefined,
      area: area || undefined,
      city: city || undefined,
      deliveryMethod: deliveryMethod && deliveryMethod !== '包邮' ? deliveryMethod : undefined,
      priceMax: priceMax > 0 ? priceMax : undefined, // 仅在大于0时发送
      priceMin: priceMin > 0 ? priceMin : undefined,
      province: province || undefined,
      publishDate: publishDate || undefined,
      shippingCost: shippingCost,
      page,
      limit,
      searchQuery: searchQuery || undefined
    }
  })
}

// 发布闲置
export const postProductAPI = (data) => {
  const userStore = useUserStore()
  return httpInstance({
    url: '/postProduct',
    method: 'post',
    headers: {
      Authorization: `${userStore.userInfo.token}`
    },
    data: data
  })
}
