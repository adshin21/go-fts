package utils

type Index map[string][]int

// make inverted index, what word is present in which document
// Eg -
// 0 - There was a bule car on the hill top.
// 1 - The bule color is top one.
// Eleminating stopword, stemmers etc the index will be
//
// bule: [0, 1]
// car: [0]
// hill: [0]
// top: [0, 1]
func (idx Index) Add(docs []document) {

	for _, doc := range docs {

		// analyze will clean up the doc, and remove unnecessary
		// things required while builing the index (search)
		for _, token := range analyze(doc.Text) {
			ids := idx[token]

			// if same id is present don't add
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

// intersection returns the set intersection between a and b.
// a and b have to be sorted in ascending order and contain no duplicates.
func Intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

// search queries the Index for the given text.
func (idx Index) Search(text string) []int {
	var r []int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = Intersection(r, ids)
			}
		} else {
			// Token doesn't exist.
			return nil
		}
	}
	return r
}
