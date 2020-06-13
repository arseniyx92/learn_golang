#include <iostream>

using namespace std;

int main(){
    int t, n, m, x, y;
    bool finish, lm;
    cin >> t;
    for (int q = 0; q < t; ++q){
        lm = false;
        cin >> n >> m;
        char a[n][m];
        for (int i = 0; i < n; ++i){
            for (int j = 0; j < m; ++j){
                a[i][j] = 0;
            }
        }
        for (int i = 0; i < n; ++i){
            for (int j = 0; j < m; ++j){
                cin >> x;
                if (x == 1){
                    for (int k = 0; k < n; ++k){
                        a[k][j] = 1;
                    }
                    for (int k = 0; k < m; ++k){
                        a[i][k] = 1;
                    }
                    a[i][j] = 2;
                }
            }
        }
        while (true){
            if (lm == true){
                lm = false;
            } else{
                lm = true;
            }
            x = y = -1;
            finish = false;
            for (int i = 0; i < n; ++i){
                if (finish) break;
                for (int j = 0; j < m; ++j){
                    if (a[i][j] == 2) break;
                    if (a[i][j] == 0){
                        x = i;
                        y = j;
                        finish = true;
                        break;
                    }
                }
            }
            if (x == -1) break;
            for (int k = 0; k < n; ++k){
                a[k][y] = 1;
            }
            for (int k = 0; k < m; ++k){
                a[x][k] = 1;
            }
            a[x][y] = 2;
        }
        if (lm == true){
            cout << "Vivek" << "\n";
        }else{
            cout << "Ashish" << "\n";
        }
    }
    return 0;
}