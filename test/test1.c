#include <stdio.h>
#include <stdlib.h>
#include <time.h>

typedef struct treap treap_t;
typedef struct pair pair_t;

struct treap
{
	int x;
	int y;
	struct treap *l[2];
};

struct pair
{
	struct treap *t;
	int d;
};

treap_t *new_treap(int x, int y)
{
	treap_t *n;

	n = (treap_t *)malloc(sizeof(treap_t));
	if (n == NULL)
	{
		return NULL;
	}
	n->x = x;
	n->y = y;
	n->l[0] = NULL;
	n->l[1] = NULL;
	return n;
}

treap_t *rotate(treap_t *t, int d)
{
	treap_t *n, *z;
	int f;

	f = !d;
	n = t->l[d];
	z = n->l[f];
	n->l[f] = t;
	t->l[d] = z;

	return n;
}

treap_t *insert(treap_t *t, int x, int y)
{
	int d;
	pair_t s[128];
	pair_t *p;
	treap_t *n;
	n = t;
	p = &s[0];
	while (n != NULL)
	{
		p->t = n;
		if (n->x > x)
		{
			p->d = 0;
			n = n->l[0];
		}
		else if (n->x < x)
		{
			p->d = 1;
			n = n->l[1];
		}
		else
		{
			return t;
		}
		p++;
	}
	n = new_treap(x, y);
	p--;
	p->t->l[p->d] = n;
	while ((p > &s[0]) && (p->t->l[p->d]->y > p->t->y))
	{
		n = p->t;
		d = p->d;
		p--;
		p->t->l[p->d] = rotate(n, d);
	}
	if ((p == &s[0]) && (p->t->l[p->d]->y > p->t->y))
	{
		return rotate(p->t, p->d);
	}
	return t;
}

treap_t *search(treap_t *t, int x)
{
	treap_t *n;

	n = t;
	while (n != NULL)
	{
		if (n->x > x)
		{
			n = n->l[0];
		}
		else if (n->x < x)
		{
			n = n->l[1];
		}
		else
		{
			return n;
		}
	}
	return NULL;
}

int main()
{
	int i, x, y, c;
	treap_t *t, *r;
	clock_t estimate;

	c = 50000000;
	t = new_treap(0, rand());
	estimate = clock();
	for (i = 1; i <= c; i++)
	{
		y = rand();
		t = insert(t, i, y);
	}

	for (i = 1; i <= c; i++)
	{
		r = search(t, i);
		if (r->x != i) {
			printf("Search failed.\n");
			break;
		}
	}

	estimate = clock() - estimate;
	printf("Time taken: %.2fs\n", ((double)estimate) / CLOCKS_PER_SEC);
	system("pause");
	return 0;
}

